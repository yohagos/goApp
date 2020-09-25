# goApp

Die auf GO basierende Applikation soll ein einfaches Telefonbuch als Web Service darstellen. Die Applikation ist via "http://localhost:5000" erreichbar. Sämtliche Testinformation können durch die zusätzlichen Erweiterungen zur URI getestet werden.

Bei der Ausführung dieser Applikation können die folgenden HTML-spezifischen Methoden ausgeführt werden:
- GET 
  - Für alle Einträge : "/contacts"
  - Für einen spezifischen Eintrag : "/contacts/{name}"
- POST
  - zum erzeugen eines Eintrags : "/contacts"
- PUT
  - Zum updaten eines Eintrags : "/contacts/{name}"
- DELETE 
  - Entfernen eines Eintrags : "/contacts/{name}"

Als Testumgebung habe ich Postman genutz, mit dem die unterschiedlichen Funktion sehr einfach zu testen waren. Drei Testdaten (Objekte) wurden mit den folgenden Namen erzeugt:
- Yosef
- Michael
- Daniel
