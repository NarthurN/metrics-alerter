package handler

const metricsTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Metrics</title>
</head>
<body>
    <h1>All Metrics</h1>
    <ul>
        {{range .}}
            <li>{{.ID}}: {{if eq .MType "gauge"}}{{.Value}}{{else}}{{.Delta}}{{end}}</li>
        {{end}}
    </ul>
</body>
</html>
`
