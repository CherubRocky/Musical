# Musical  
## Author: Zaldivar Alanis Rodrigo
## Requirements
Se necesita tener la versión 1.23.1 de ___Go___ instalada.
Se puede obtener de su stitio web oficial: [Golang](https://go.dev/)
También se necesita tener instalado ___sqlite3___, que se puede obtener desde el
gestor de paquetes de linux. O desde su página web oficial: [sqlite3](https://www.sqlite.org/)


## Installation
Para instalar:
Ejecuta 'git clone https://github.com/CherubRocky/Musical.git' en la carpeta de tu gusto.
Muévete al directorio raíz del proyecto
ejecuta: "go mod tidy"
luego ejecuta "go build"
finalmente ejecuta "./Musical"

## Uso del buscador
Para usar la búsqueda, introduce el nombre de la canción deseada luego haz click en buscar.
Al hacer esto, se buscara el nombre de la canción textualmente.

Para hacer una búsqueda más precisa, puedes hacer lo siguiente:
Puedes buscar por intérprete, y/o título y/o álbum.
para buscar por intérprete escribe en el campo "p:<nombre del intérprete>"
para buscar por álbum escribe "a:<nombre del álbum>"
para buscar por título escribe "t:<título de la canción>"
Al usar esta notación, se buscan las canciones que tengan contenido lo que sigue después de ':'
(es un poco más laxo).

Si quieres buscar las canciones que cumplan con 3 o 2 de estas características al mismo tiempo:
pon los campos que quieras que cumplan usando la notación anterior, pero sepáralos por comas.

Ejemplo1: Si quieres buscar todas las canciones del álbum grandes Éxitos de José José, escribe "p:José José, a:Grandes Éxitos"
Ejemplo2: Si quieres buscar la canción Nubes del album El Silencio de la banda Caifanes, escribe
"t:Nubes,p:Caifanes, a:El Silencio"


Si quieres listar todas las canciones introduce el caracter '*'.


## Uso del minero
Para minar las canciones de tu directorio "~/Música" o "~/Music", solo es necesario hacer click en el botón "importar canciones".
