package adapter

import (
	"bytes"
	"html/template"
	"time"

	"github.com/isaias-dgr/story-balance/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

// TODO crear un mail template en storage
const templateMail string = `<!DOCTYPE html>
<html>

<head>
    <title>Balance de Cuenta</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }

        .container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
        }

        h1 {
            font-size: 24px;
            margin-top: 0;
            margin-bottom: 20px;
        }

        table {
            width: 100%;
            border-collapse: collapse;
        }

        th,
        td {
            padding: 10px;
            text-align: left;
            border-bottom: 1px solid #ddd;
        }

        th {
            background-color: #f5f5f5;
            font-weight: normal;
        }

        .balance {
            font-size: 18px;
            font-weight: bold;
        }
    </style>
</head>

<body>
    <header class="container">
        <h1>Balance de Cuenta</h1>
        <span>User: {{.Id}}</span>
        <span>Peridod: {{.Year}} {{.Month}}</span>
    </header>
    {{range .Products}}
    <div class="container">
        <h1>{{.Id}}</h1>
        <div>
            Total {{ printf "$%.2f" .Total}}
        </div>
        <table>
            <thead>
                <tr>
                    <th>Fecha</th>
                    <th>Descripción</th>
                    <th>Monto</th>
                </tr>
            </thead>
            {{range .Transactions}}
            <tbody>
                <tr>
                    <td>{{.Date | formatDate}}</td>
                    <td>{{.Description}} <span>{{.Code}}</span></td>
                    <td>{{.Amount}}</td>
                </tr>
            </tbody>
            {{end}}
            <tfoot>
                <tr>
                    <th colspan="2">Balance total:</th>
                    <td class="balance">{{.Total}}</td>
                </tr>
            </tfoot>
        </table>
    </div>
    {{end}}
</body>

</html>`

type Writter struct {
	log      *log.Logger
	template string
}

func NewWritter(l *log.Logger) *Writter {
	return &Writter{
		log:      l,
		template: templateMail,
	}
}

func (m Writter) GetDoc(acc *domain.Account) (*bytes.Buffer, error) {
	funcMap := template.FuncMap{
		"formatDate": func(date time.Time) string {
			return date.Format("01/02")
		},
	}
	tmpl := template.Must(template.New("list").Funcs(funcMap).Parse(m.template))
	buffer := new(bytes.Buffer)
	err := tmpl.Execute(buffer, acc)
	if err != nil {
		return nil, err
	}
	m.log.Info("Document created!")
	return buffer, nil
}
