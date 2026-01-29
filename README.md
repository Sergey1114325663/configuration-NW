# Network Configuration Manager

Инструмент для управления VLAN, VPN, Firewall, QoS и мониторингом трафика.

## Возможности

### 1. VLAN Management
- Создание виртуальных локальных сетей
- Управление интерфейсами VLAN
- Просмотр списка VLAN
- Удаление VLAN

```bash
./network-config vlan create --id 10 --name Production --description "Production Network"
./network-config vlan list
./network-config vlan delete --id 10
```

### 2. VPN Configuration

#### L2TP (Layer 2 Tunneling Protocol)
```bash
./network-config vpn l2tp --action create --name office-vpn
./network-config vpn l2tp --action connect --name office-vpn
./network-config vpn l2tp --action disconnect --name office-vpn
```

#### PPTP (Point-to-Point Tunneling Protocol)
```bash
./network-config vpn pptp --action create --name remote-vpn
./network-config vpn pptp --action connect --name remote-vpn
```

#### SSTP (Secure Socket Tunneling Protocol)
```bash
./network-config vpn sstp --action create --name secure-vpn
./network-config vpn sstp --action connect --name secure-vpn
```

#### OpenVPN
```bash
./network-config vpn openvpn --action start --config /path/to/config.ovpn
./network-config vpn openvpn --action stop --config /path/to/config.ovpn
```

### 3. Firewall Management
```bash
./network-config firewall rule --action add --rule "allow tcp port 80"
./network-config firewall status
./network-config firewall list
```

### 4. Quality of Service (QoS)
```bash
./network-config qos policy --name high-priority --bandwidth 100Mbps
./network-config qos limit --interface eth0 --limit 50Mbps
./network-config qos priority --class video --priority 5
```

### 5. Traffic Monitoring
```bash
./network-config monitor traffic --interface eth0
./network-config monitor stats
./network-config monitor alerts --threshold 100Mbps
```

## Структура проекта

```
network-config/
├── main.go           # Точка входа приложения
├── go.mod           # Go модули
├── README.md        # Документация
├── config/
│   └── config.go    # Структуры конфигурации
├── cmd/
│   └── commands.go  # CLI команды
└── examples/
    ├── network.yml  # Пример конфигурации сети
    └── vpn.yml      # Пример конфигурации VPN
```

## Установка

```bash
go mod download
go build -o network-config main.go
```

## Использование

### Основные команды

```bash
# Просмотр справки
./network-config --help

# Справка по подкоманде
./network-config vlan --help
./network-config vpn --help
./network-config firewall --help
./network-config qos --help
./network-config monitor --help
```

## Примеры конфигурации

### network.yml
```yaml
vlans:
  - id: "1"
    name: "Management"
    description: "Management VLAN"
    interfaces: ["eth0"]
    status: "active"
  - id: "10"
    name: "Production"
    description: "Production Network"
    interfaces: ["eth1", "eth2"]
    status: "active"

vpn:
  - name: "office-vpn"
    type: "L2TP"
    server: "vpn.example.com"
    port: 1194
    username: "user"
    password: "pass"
    enabled: true

firewall:
  - id: "1"
    direction: "Inbound"
    action: "Allow"
    protocol: "TCP"
    port: "80"
  - id: "2"
    direction: "Inbound"
    action: "Allow"
    protocol: "TCP"
    port: "443"

qos:
  - name: "high-priority"
    interface: "eth0"
    bandwidth: "100Mbps"
    priority: 7

monitoring:
  enabled: true
  interval: 5
  logfile: "/var/log/network-monitor.log"
```

## Зависимости

- `github.com/sirupsen/logrus` - логирование
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - конфигурация
- `gopkg.in/yaml.v3` - YAML парсинг

## Лицензия

MIT
