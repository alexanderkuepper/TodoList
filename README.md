# Ohne Docker 

Mit `go install` lassen sich alle Abhängigkeiten installieren. Dieser Befehl muss im Verzeichnis `todolist` ausgeführt werden. Im Anschluss lässt sich das Projekt mit `go run main.go` im selben Verzeichnis starten.  

# Mit Docker 

Mit `docker build --tag todolist .` lässt sich der Container bauen. Dieser Befehl muss im Verzeichnis `todolist` ausgeführt werden. 

Im Anschluss kann der Container mit `docker run --publish 8080:8080 todolist` gestartet werden. 

# Doku 

Die API Dokumentation findet man unter http://localhost:8080/swagger/index.html