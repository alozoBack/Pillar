package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"github.com/joho/godotenv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh" // Используем правильный пакет SSH
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

const (
	host    = "0.0.0.0"
	port    = "22"
	keyPath = ".ssh/host_key" // Путь к SSH ключу
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

// Генерация SSH ключей, если их нет
func generateHostKey() {
	// Проверка, существует ли директория .ssh, и создание её при необходимости
	if _, err := os.Stat(".ssh"); os.IsNotExist(err) {
		err = os.Mkdir(".ssh", 0700) // Создаем директорию с правами доступа
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create .ssh directory:", err)
			os.Exit(1)
		}
	}

	// Генерация ключа, если он не существует
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		err := exec.Command("ssh-keygen", "-t", "rsa", "-f", keyPath, "-N", "").Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to generate host key:", err)
			os.Exit(1)
		}
	}
}

// teaHandlerFunc обрабатывает каждую SSH-сессию и создает Bubble Tea программу для каждого пользователя
func teaHandlerFunc(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	m := NewModel()                 // Создаем новую модель для каждого сеанса
	return m, []tea.ProgramOption{} // Возвращаем модель и параметры программы (если есть)
}

// Стартуем SSH сервер с Bubble Tea
func startSSHServer() {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(keyPath),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandlerFunc),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}

}

func main() {
	// Генерация SSH ключей, если их нет
	generateHostKey()

	// Запуск SSH сервера
	startSSHServer()
}
