# Milestones del proyecto

## [[M0] Modelo del problema](https://github.com/GaelGoncalba/AutoShopping/milestone/1)
### Descripción: 
Representación del modelo del problema concretado en la HU02 (Fernando).

### Entregable: 
- Código que represente el modelo del problema, sobre el cual se desarrollará posteriormente.

### Viabilidad: 
Este milestone se considerará viable si:
- El código entregado representa el modelo del problema abarcando los problemas descritos en la HU02.

----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M1] Generación automática de una lista de la compra](https://github.com/GaelGoncalba/AutoShopping/milestone/2)
### Descripción: 
Desarrollo de un módulo que extraiga los datos de las fuentes seleccionadas para generar automáticamente una lista de la compra que contenga todos los alimentos necesarios para satisfacer las necesidades nutricionales de una persona adulta durante 7 días, asegurando que la dieta sea lo más equilibrada posible, atendiendo al problema descrito en la HU02 (Fernando).

### Entregable: 
Código que implemente:
- Módulo funcional que extraiga los datos necesarios de productos de supermercado desde la web [SoySuper](https://soysuper.com/).
- Módulo funcional que genera automáticamente una lista de la compra a partir de los datos extraídos.
- Documentación mínima en el propio código que describa los criterios de balance nutricional empleados y cómo se aplican en la generación de la lista.

### Viabilidad: 
Este milestone se considerará viable cuando:
- La lógica de selección sea evaluada mediante pruebas nutricionales básicas, asegurando que los resultados cumplen con los mínimos recomendados para una dieta saludable (por ejemplo, equilibrio de macronutrientes).

----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M2] Ajuste y filtrado de la lista de la compra a un presupuesto y preferencias](https://github.com/GaelGoncalba/AutoShopping/milestone/3)
### Descripción: 
Implementación de un módulo que genere una lista de la compra ajustada a un presupuesto semanal ingresado por el usuario, aplicando además filtros de preferencias que permitan seleccionar tipos de alimentos (pescado, verdura, carne, etc.) y supermercados específicos. El sistema debe excluir productos que no cumplan con los filtros definidos sin comprometer el equilibrio nutricional, resolviendo los problemas descritos en la HU01 (Julián) y la HU03 (Nieves).

### Entregable:
Código que implemente:
- Módulo funcional que genera una lista de la compra ajustada a un presupuesto.
- Módulo que filtre las listas de la compra según las preferencias del usuario en cuanto a tipos de productos y supermercados específicos.

### Viabilidad:
Este milestone se considerará viable si:
- Las listas generadas respetan al 100% el presupuesto máximo ingresado y las preferencias del usuario.
- El sistema es capaz de ofrecer combinaciones de productos variadas y ajustadas a las restricciones sin comprometer la calidad nutricional.
