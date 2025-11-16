# Rocket-Lab

![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/mercuryqa/f86f7beb53796d3daf361c41dfb12e85/raw/coverage.json)

Для того чтобы вызывать команды из Taskfile, необходимо установить Taskfile CLI:

```bash
brew install go-task
```

## CI/CD

Проект использует GitHub Actions для непрерывной интеграции и доставки. Основные workflow:

- **CI** (`.github/workflows/ci.yml`) - проверяет код при каждом push и pull request
    - Линтинг кода
    - Проверка безопасности
    - Выполняется автоматическое извлечение версий из Taskfile.yml
