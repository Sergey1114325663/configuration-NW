# Network Configuration Manager - API Documentation

## REST API Endpoints

### VLAN Management

#### Get all VLANs
```
GET /api/v1/vlans
Response: 200 OK
[
  {
    "id": "1",
    "name": "Management",
    "status": "active",
    "members": ["eth0"],
    "mtu": 1500
  }
]
```

#### Create VLAN
```
POST /api/v1/vlans
Content-Type: application/json
{
  "id": "10",
  "name": "Production",
  "status": "active",
  "members": ["eth1"],
  "mtu": 1500
}
Response: 201 Created
```

#### Get specific VLAN
```
GET /api/v1/vlans/{id}
Response: 200 OK
{
  "id": "10",
  "name": "Production",
  "status": "active"
}
```

#### Delete VLAN
```
DELETE /api/v1/vlans/{id}
Response: 204 No Content
```

### VPN Management

#### Get all VPN connections
```
GET /api/v1/vpn/connections
Response: 200 OK
[
  {
    "name": "office-vpn",
    "type": "L2TP",
    "status": "connected",
    "ip_address": "10.0.0.5"
  }
]
```

#### Create VPN connection
```
POST /api/v1/vpn/connections
Content-Type: application/json
{
  "name": "office-vpn",
  "type": "L2TP",
  "server": "vpn.example.com",
  "port": 1194,
  "username": "user",
  "password": "pass"
}
Response: 201 Created
```

#### Connect to VPN
```
POST /api/v1/vpn/connections/{name}/connect
Response: 200 OK
{ "status": "connected" }
```

#### Disconnect VPN
```
POST /api/v1/vpn/connections/{name}/disconnect
Response: 200 OK
{ "status": "disconnected" }
```

### Firewall Management

#### Get firewall rules
```
GET /api/v1/firewall/rules
Response: 200 OK
[
  {
    "id": "1",
    "direction": "inbound",
    "action": "allow",
    "protocol": "tcp",
    "port": "80"
  }
]
```

#### Add firewall rule
```
POST /api/v1/firewall/rules
Content-Type: application/json
{
  "id": "1",
  "direction": "inbound",
  "action": "allow",
  "protocol": "tcp",
  "port": "80",
  "priority": 10
}
Response: 201 Created
```

#### Get firewall status
```
GET /api/v1/firewall/status
Response: 200 OK
{
  "status": "enabled",
  "active_rules": 42,
  "blocked_packets": 1234
}
```

#### Delete rule
```
DELETE /api/v1/firewall/rules/{id}
Response: 204 No Content
```

### QoS Management

#### Get QoS policies
```
GET /api/v1/qos/policies
Response: 200 OK
[
  {
    "name": "high-priority",
    "interface": "eth0",
    "bandwidth": "100Mbps",
    "priority": 7
  }
]
```

#### Create QoS policy
```
POST /api/v1/qos/policies
Content-Type: application/json
{
  "name": "high-priority",
  "interface": "eth0",
  "bandwidth": "100Mbps",
  "priority": 7,
  "cir": 100000,
  "pir": 150000
}
Response: 201 Created
```

#### Update QoS policy
```
PUT /api/v1/qos/policies/{name}
Content-Type: application/json
{
  "bandwidth": "120Mbps",
  "priority": 8
}
Response: 200 OK
```

#### Delete QoS policy
```
DELETE /api/v1/qos/policies/{name}
Response: 204 No Content
```

### Traffic Monitoring

#### Get real-time traffic
```
GET /api/v1/monitor/traffic?interface=eth0
Response: 200 OK
{
  "interface": "eth0",
  "rx_bytes": 1234567,
  "tx_bytes": 987654,
  "rx_packets": 5432,
  "tx_packets": 3210,
  "rx_errors": 0,
  "tx_errors": 0,
  "bandwidth_rx": 80.5,
  "bandwidth_tx": 40.2
}
```

#### Get network statistics
```
GET /api/v1/monitor/stats
Response: 200 OK
{
  "total_traffic": 2.2,
  "avg_bandwidth": 50,
  "peak_bandwidth": 150,
  "active_connections": 234
}
```

#### Configure traffic alerts
```
POST /api/v1/monitor/alerts
Content-Type: application/json
{
  "threshold": "100Mbps",
  "condition": "greater_than",
  "notification": "email"
}
Response: 201 Created
```

#### Get alerts
```
GET /api/v1/monitor/alerts
Response: 200 OK
[
  {
    "id": "alert-1",
    "threshold": "100Mbps",
    "triggered": true,
    "timestamp": "2026-01-29T10:30:00Z"
  }
]
```

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request",
  "message": "Missing required field: name"
}
```

### 401 Unauthorized
```json
{
  "error": "Unauthorized",
  "message": "Authentication token is required"
}
```

### 404 Not Found
```json
{
  "error": "Not found",
  "message": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error",
  "message": "An unexpected error occurred"
}
```

## Authentication

All API requests (except health check) require authentication token:

```
Authorization: Bearer <token>
```

### Get authentication token
```
POST /api/v1/auth/token
Content-Type: application/json
{
  "username": "admin",
  "password": "password"
}
Response: 200 OK
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_in": 3600
}
```

## Rate Limiting

API rate limits:
- 1000 requests per minute per IP address
- 100 concurrent connections

Headers:
- `X-RateLimit-Limit: 1000`
- `X-RateLimit-Remaining: 999`
- `X-RateLimit-Reset: 1234567890`

## WebSocket Support

Real-time traffic monitoring:

```javascript
ws = new WebSocket("wss://api.example.com/api/v1/monitor/traffic/ws?interface=eth0");

ws.onmessage = function(event) {
  console.log(JSON.parse(event.data));
};

// Example output:
// {
//   "interface": "eth0",
//   "rx_bytes": 1234567,
//   "tx_bytes": 987654,
//   "timestamp": "2026-01-29T10:30:00Z"
// }
```

## Webhooks

Configure webhooks for events:

```
POST /api/v1/webhooks
{
  "event": "firewall_rule_blocked",
  "url": "https://example.com/webhook",
  "secret": "webhook-secret"
}
```

## CLI Integration

Use curl to interact with the API:

```bash
# Get all VLANs
curl -H "Authorization: Bearer TOKEN" http://localhost:8080/api/v1/vlans

# Create VLAN
curl -X POST http://localhost:8080/api/v1/vlans \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"id":"10","name":"Production"}'

# Monitor traffic
curl http://localhost:8080/api/v1/monitor/traffic?interface=eth0
```
