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