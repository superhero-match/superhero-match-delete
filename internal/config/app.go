/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package config

// App holds the configuration values for the application.
type App struct {
	Port       string `env:"SUPERHERO_MATCH_DELETE_APP_PORT" yaml:"port" default:":4300"`
	CertFile   string `env:"SUPERHERO_MATCH_DELETE_APP_CERT_FILE" yaml:"cert_file" default:"./cmd/api/certificate.pem"`
	KeyFile    string `env:"SUPERHERO_MATCH_DELETE_APP_KEY_FILE" yaml:"key_file" default:"./cmd/api/key.pem"`
	TimeFormat string `env:"SUPERHERO_MATCH_DELETE_APP_TIME_FORMAT" yaml:"time_format" default:"2006-01-02T15:04:05"`
}
