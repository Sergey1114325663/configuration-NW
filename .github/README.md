# GitHub Configuration Files

## Структура
```
.github/
├── workflows/
│   └── go-build-test.yml (CI/CD pipeline)
├── ISSUE_TEMPLATE/
│   ├── bug_report.md
│   └── feature_request.md
└── pull_request_template.md
```

## Использование

### Создание Новой Issue
GitHub автоматически предложит выбрать из доступных templates:
- Bug Report (для ошибок)
- Feature Request (для новых функций)

### Создание Pull Request
При создании PR автоматически загружается `pull_request_template.md` с основным структурой для описания изменений.

## Автоматизация

Workflow файл в `.github/workflows/go-build-test.yml` автоматически:
- Запускает тесты при каждом commit/PR
- Проверяет код через golangci-lint
- Собирает бинарники для Windows
- Собирает отчет о покрытии (coverage)

## Настройки Branch Protection (рекомендуется для GitHub)

После первого push, рекомендуется включить:

1. **Require pull request reviews before merging**
   - Settings → Branches → Branch protection rules
   - Require 1 review minimum

2. **Require status checks to pass before merging**
   - Require Go build and tests to pass

3. **Require branches to be up to date before merging**
   - Ensure develop is in sync before merge

4. **Include administrators** 
   - Applies rules to everyone including admins

## Рекомендуемые GitHub Actions

```yaml
# Дополнительно рекомендуется:
- Code scanning (Dependabot)
- Security alerts
- Auto-merge для minor version updates
```
