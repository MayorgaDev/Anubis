package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ProcessEmail(email string) (string, error) {
	if err := validateNotEmpty(email); err != nil {
		return "", err
	}

	email = strings.TrimSpace(email)

	if err := validateMaxLength(email); err != nil {
		return "", err
	}

	if err := validateFormat(email); err != nil {
		return "", err
	}

	return sanitizeEmail(email), nil
}

func validateNotEmpty(email string) error {
	if email == "" {
		return fmt.Errorf("Ingrese un correo electronico.")
	}
	return nil
}

func validateMaxLength(email string) error {
	if len(email) > 150 {
		return fmt.Errorf("El correo electronico es demasiado largo, debe tener menos de 150 caracteres.")
	}
	return nil
}

func validateFormat(email string) error {
	emailRegex := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return fmt.Errorf("El formato del correo electrónico no es válido.")
	}
	return nil
}

func sanitizeEmail(email string) string {
	email = strings.ReplaceAll(email, "<", "&lt;")
	email = strings.ReplaceAll(email, ">", "&gt;")
	email = strings.ReplaceAll(email, "'", "&#039;")
	email = strings.ReplaceAll(email, "\"", "&quot;")
	return email
}

func main() {
	var correo string
	fmt.Print("Introduce tu coreo: ")
	fmt.Scanln(&correo)

	email, err := ProcessEmail(correo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El correo electronico: ", email, " es valido.")
}
