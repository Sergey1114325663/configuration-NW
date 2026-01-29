# Network Configuration - Advanced Setup

## VLAN Configuration

### 802.1Q Tagging
```
Interface: eth0.1 → VLAN 1 (Management)
Interface: eth0.10 → VLAN 10 (Production)
Interface: eth0.20 → VLAN 20 (Development)
```

## VPN Protocols Configuration

### L2TP/IPSec
- Порт: UDP 500 (IKE), UDP 4500 (IPSec), UDP 1701 (L2TP)
- Шифрование: AES-256
- Аутентификация: SHA-256
- MTU: 1500 bytes

### PPTP
- Порт: TCP 1723 (контроль), GRE 47 (данные)
- Шифрование: MPPE 128-bit
- Аутентификация: PAP, CHAP
- Для Windows клиентов

### SSTP
- Порт: TCP 443 (HTTPS)
- Шифрование: SSL/TLS + PPP
- Работает через HTTP прокси
- Поддержка NAT Traversal

### OpenVPN
- Порт: UDP 1194 (рекомендуется)
- Шифрование: AES-256-CBC
- Аутентификация: HMAC-SHA256
- Кросс-платформенное решение

## Firewall Rules Priority

```
Priority 1: Management Traffic (SSH, RDP)
Priority 2: VPN Traffic
Priority 3: Web Traffic (HTTP/HTTPS)
Priority 4: DNS Traffic
Priority 5: Other Traffic
Priority 6: Deny Rules
```

## QoS Scheduling Algorithms

1. **FIFO** - First In First Out
2. **RR** - Round Robin (для высокого приоритета)
3. **WFQ** - Weighted Fair Queuing (стандартный)
4. **Priority Queue** - Для критичного трафика

## Traffic Monitoring Metrics

```
- RX/TX bytes and packets
- Error rate
- Packet loss
- Latency
- Bandwidth utilization
- Connection count
- Protocol distribution
```

## Performance Optimization

### VLAN Optimization
- Минимизировать inter-VLAN traffic
- Использовать VLAN ACLs для фильтрации
- Регулярная аудит VLAN конфигурации

### VPN Optimization
- Выбор оптимального протокола
- Сжатие данных (LZO)
- Fragment/MTU adjustment
- Load balancing между VPN серверами

### Firewall Optimization
- Порядок правил (часто используемые сверху)
- Использование groups и aliases
- Регулярный audit неиспользуемых правил
- State tracking для SYN protection

### QoS Optimization
- Правильная классификация трафика
- Балансировка полосы пропускания
- Adaptive QoS для переменного трафика
- Постоянный мониторинг

## Security Best Practices

1. **Authentication**
   - Использовать VPN с сертификатами
   - Двухфакторная аутентификация
   - Регулярная смена паролей

2. **Encryption**
   - AES-256 как минимум
   - TLS 1.2+ для SSL/TLS
   - Perfect Forward Secrecy (PFS)

3. **Access Control**
   - Principle of Least Privilege
   - Регулярный audit доступа
   - Time-based access restrictions

4. **Monitoring**
   - Логирование всех сетевых событий
   - Алерты на аномальный трафик
   - DDoS protection
   - Intrusion Detection System (IDS)

## High Availability Setup

### VPN HA
```
VPN Server 1 (Primary) ← → Active
VPN Server 2 (Secondary) ← → Standby
```

### Firewall HA
```
Firewall 1 (Active) → 99.9% uptime
Firewall 2 (Passive) → Failover
```

### QoS HA
- Load balancing между несколькими узлами
- Automatic failover на резервные каналы
