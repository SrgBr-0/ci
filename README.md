# Локальный запуск GitHub Actions с использованием `act`

Этот документ описывает, как настроить и запустить GitHub Actions на локальной машине с помощью инструмента `act`.

## Установка

### 1. Установка `act`
Для установки `act` используйте Homebrew(MacOS/Linux):
```bash
brew install act
```
Через curl (все платформы):
```bash
curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
```
Для Windows:
Установите act через сборку с [релиза](https://github.com/nektos/act/releases).

---
Вам нужно выбрать образ, который будет использоваться для выполнения GitHub Actions локально.

#### Объяснение вариантов:
##### Large: Полный образ с поддержкой всех инструментов, которые доступны в GitHub-hosted runners.
 - Плюсы: Поддерживает практически все действия, используемые в workflows.
 - Минусы: Требует ~75GB свободного пространства (17GB для скачивания + 53.1GB для хранения).
##### Medium (рекомендуется): Уменьшенный образ, включающий только необходимые инструменты для выполнения большинства Actions.
 - Плюсы: Компактнее (около 500MB), быстрее скачивается.
 - Минусы: Может не поддерживать редкие действия.
##### Micro: Минимальный образ, содержащий только Node.js.
 - Плюсы: Самый маленький (менее 200MB).
 - Минусы: Поддерживает только базовые Node.js workflows, не подходит для Docker-based задач.
---
### 2. Проверка установки Docker
```bash
docker version
```
---
### 3. Запуск GitHub Actions локально
   #### Перейдите в корневую директорию вашего проекта, где находится файл конфигурации GitHub Actions (.github/workflows/).
   #### Запустите act:
```bash
act
```