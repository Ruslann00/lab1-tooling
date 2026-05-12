# Lab 1 Tooling

Лабораторна робота №1: Налаштування професійного інструментарію розробки на Go.

## Опис

Проект демонструє базове налаштування Go-розробки:

- Структура Go-проекту;
- unit-тести;
- table-driven tests;
- Статичний аналіз через golangci-lint;
- автоматизація через Makefile;
- Складання бінарного файлу.

## Структура проекту

```text
.
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── calculator.go
│   └── calculator_test.go
├── .golangci.yml
├── .gitignore
├── Makefile
├── README.md
├── go.mod


## Pull Request

Ця гілка створена для оформлення Pull Request до лабораторної роботи №1.

# Lab 1 Tooling / Lab 2 Dependencies and CI

![Go CI](https://github.com/Ruslann00/lab1-tooling/actions/workflows/ci.yml/badge.svg)

## Лабораторна робота №2

Тема: Управління залежностями в Go та автоматизація перевірок через GitHub Actions.

## Використані залежності

У проєкті додано:

- `go.uber.org/zap` — бібліотека для логування;
- `github.com/spf13/viper` — бібліотека для читання конфігураційного файлу.

## Конфігурація

Файл `config.yaml`:

```yaml
app:
  name: "Lab 2 Go Tooling"
  environment: "development"

calculator:
  a: 10
  b: 5