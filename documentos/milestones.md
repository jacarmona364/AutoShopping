# Milestones del proyecto

## [[M0] Modelo del problema](https://github.com/GaelGoncalba/AutoShopping/milestone/1)
### Descripción: 
Analizar las historias de usuario HU01 (Julián) y HU02 (Fernando) para desarrollar el modelo del problema.

### Entregables: 
- Código que concrete el modelo del problema, sobre el cual se desarrollará posteriormente.
  
### Producto: 
- Código que define el modelo de la aplicación.

### Viabilidad: 
Este milestone se considerará viable si:
- El código entregado representa el modelo del problema abarcando los problemas descritos en las HU01 y HU02.
- El código está correctamente estructurado, permitiendo el futuro desarrollo de la aplicación.

### Historias de usuario asociadas:
- HU01 (Julián) y HU02 (Fernando)

----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M1] Generación automática de una lista de la compra](https://github.com/GaelGoncalba/AutoShopping/milestone/2)
### Descripción: 
Desarrollar un módulo que extraiga los datos de las fuentes seleccionadas para generar automáticamente una lista de la compra que contenga todos los alimentos necesarios para satisfacer las necesidades nutricionales de una persona adulta durante 7 días, asegurando que la dieta sea lo más equilibrada posible.

### Entregables: 
- Módulo funcional que extraiga los datos necesarios de productos de supermercado desde la web [SoySuper](https://soysuper.com/).
- Módulo funcional que genera automáticamente una lista de la compra a partir de los datos extraídos.
- Documentación mínima en el propio código que describa los criterios de balance nutricional empleados y cómo se aplican en la generación de la lista.

### Producto: 
- Código del módulo base de la aplicación que extrae los datos necesarios de la web seleccionada y genera una lista de la compra sin restricciones adicionales (en concreto, ajustes según presupuesto o preferencias dietéticas) de acuerdo con los valores recomendados para una dieta equilibrada.

### Viabilidad: 
Este milestone es fundamental para validar que el módulo de generación de listas funciona correctamente y es capaz de producir resultados válidos. Se debe tener en cuenta:
- La lógica de selección debe ser evaluada mediante pruebas nutricionales básicas, asegurando que los resultados cumplen con los mínimos recomendados para una dieta saludable (por ejemplo, equilibrio de macronutrientes).
- La extracción de datos desde la web seleccionada es completa y satisfactoria, obteniendo la información necesaria para completar los elementos del dominio de la aplicación.
- Asegurar que la lógica del módulo se puede adaptar fácilmente a futuras mejoras y restricciones.

### Historias de usuario asociadas:
- HU02 (Fernando)

----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M2] Ajuste y filtrado de la lista de la compra a un presupuesto y preferencias](https://github.com/GaelGoncalba/AutoShopping/milestone/3)
### Descripción: 
Implementar un módulo que genere una lista de la compra ajustada a un presupuesto semanal ingresado por el usuario, aplicando además filtros de preferencias que permitan seleccionar tipos de alimentos (pescado, verdura, carne, etc.) y supermercados específicos. El sistema debe excluir productos que no cumplan con los filtros definidos sin comprometer el equilibrio nutricional.

### Entregables:
- Módulo funcional que genera una lista de la compra ajustada a un presupuesto, con productos filtrados por tipo de alimento y supermercado.
- Código que integre tanto las restricciones de presupuesto como las preferencias del usuario, priorizando productos más accesibles, manteniendo una dieta equilibrada.

### Producto:
- Funcionalidad que permite al usuario ingresar un presupuesto semanal, aplicar filtros de tipo de alimento y seleccionar supermercados específicos. La lista generada deberá respetar el presupuesto y las preferencias indicadas, actualizando el módulo de generación automática de listas de la compra implementado en el milestone 1. El algoritmo funcionará priorizando los productos más accesibles y necesarios dentro del presupuesto, excluyendo aquellos que no cumplan con las restricciones impuestas por el usuario. El sistema realizará un análisis iterativo para seleccionar los productos que maximicen el valor nutricional mientras se mantiene dentro del límite de costo.

### Viabilidad:
Este milestone asegurará la correcta personalización de la generación de listas de la compra por parte del usuario. Se deberá asegurar que:
- Las listas generadas respetan al 100% el presupuesto máximo ingresado y las preferencias del usuario.
- El sistema es capaz de ofrecer combinaciones de productos variadas y ajustadas a las restricciones, sin comprometer la calidad nutricional.

### Historias de usuario asociadas:
- HU01 (Julián) y HU03 (Nieves)
