
# ApiRAE

API para la Real Academia de la lengua Española.
Extrae las definiciones de la web de la RAE.

## Forma de uso:
    `https://apirae.herokuapp.com/buscar/[termino de búsqueda]`

## Devuelve json:
    {
        acepcion: acepción,
        definiciones : {
            id: identificador único,
            definicion: significado
        }
    }

## Ejemplo:
    `https://mi-url/buscar/perro` 

[
  {
    "Title": "Del lat. casa 'choza'.",
    "Data": [
      {
        "id": 1,
        "data": "1. f. Edificio para habitar. Una casa de ocho plantas."
      },
      {
        "id": 2,
        "data": "2. f. Edificio de una o pocas plantas destinado a vivienda unifamiliar, en oposición a piso. Quieren vender el piso y comprarse una casa."
      },
      {
        "id": 3,
        "data": "3. f. piso (‖ vivienda). Mi casa está en el 3.º C."
        ....
