package master_ui

import (
	"html/template"
)

var StatusTpl = template.Must(template.New("status").Parse(`<!DOCTYPE html>
<html>
  <head>
    <title>SeaweedFS {{ .Version }}</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css">
  </head>
  <body>
    <div class="container">
      <div class="page-header">
        <h1>SeaweedFS <small>{{ .Version }}</small></h1>
      </div>

      <div class="row">
        <div class="col-sm-6">
          <h2>Cluster status</h2>
          <table class="table">
            <tbody>
              <tr>
                <th>Free</th>
                <td>{{ .Topology.Free }}</td>
              </tr>
              <tr>
                <th>Max</th>
                <td>{{ .Topology.Max }}</td>
              </tr>
              <tr>
                <th>Leader</th>
                <td><a href="http://{{ .Leader }}">{{ .Leader }}</a></td>
              </tr>
              <tr>
                <td class="col-sm-2 field-label"><label>Peers:</label></td>
                <td class="col-sm-10"><ul class="list-unstyled">
                {{ range $k, $p := .Peers }}
                  <li><a href="{{ $p.ConnectionString }}">{{ $p.Name }}</a></li>
                {{ end }}
                </ul></td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="col-sm-6">
          <h2>System Stats</h2>
          <table class="table table-condensed table-striped">
            <tr>
              <th>Concurrent Connections</th>
              <td>{{ .Counters.Connections.WeekCounter.Count }}</td>
            </tr>
          {{ range $key, $val := .Stats }}
            <tr>
              <th>{{ $key }}</th>
              <td>{{ $val }}</td>
            </tr>
          {{ end }}
          </table>
        </div>
      </div>

      <div class="row">
        <h2>Topology</h2>
        <table class="table table-striped">
          <thead>
            <tr>
              <th>Data Center</th>
              <th>Rack</th>
              <th>RemoteAddr</th>
              <th>#Volumes</th>
              <th>Max</th>
            </tr>
          </thead>
          <tbody>
          {{ range $dc_index, $dc := .Topology.DataCenters }}
            {{ range $rack_index, $rack := $dc.Racks }}
              {{ range $dn_index, $dn := $rack.DataNodes }}
            <tr>
              <td><code>{{ $dc.Id }}</code></td>
              <td>{{ $rack.Id }}</td>
              <td><a href="http://{{ $dn.Url }}/ui/index.html">{{ $dn.Url }}</a></td>
              <td>{{ $dn.Volumes }}</td>
              <td>{{ $dn.Max }}</td>
            </tr>
              {{ end }}
            {{ end }}
          {{ end }}
          </tbody>
        </table>
      </div>

    </div>
  </body>
</html>
`))
