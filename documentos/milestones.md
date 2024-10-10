# Milestones del proyecto

## [[M01] Fundamentos del problema](https://github.com/GaelGoncalba/AutoShopping/milestone/1)
### Descripción: Analizar las historias de usuario para modelar la extracción de datos (scrapping) de la web seleccionada y generar una interfaz mínima básica para las posteriores funciones que va a poder realizar el usuario.

### Entregables: 
- Código funcional del núcleo de la aplicación que implemente la extracción básica de datos y la interfaz mínima con la que posteriormente interaccionará el usuario.

### Producto: 
- Prototipo de la aplicación con una interfaz mínima que permita mostrar cómo se procesarían los datos de productos alimenticios obtenidos desde la web [SoySuper](https://soysuper.com/).

### Viabilidad: 
  Este milestone es crucial para establecer las bases tecnológicas del proyecto. Se deberá asegurar que:
- Las tecnologías seleccionadas permiten la extracción completa de la información relevante de la web analizada.
- El código generado para la interfaz permitirá en un futuro la correcta interacción del usuario con la aplicación.

### Criterios de viabilidad: 
- El prototipo debe demostrar que se han extraído correctamente todos los datos de los productos necesarios.
- La interfaz debe ser mínimamente intuitiva y permitir la implementación de la funcionalidad de _"Generación automática de lista de la compra"_.

### HU relacionadas: [HU01],[HU02],[HU03]
----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M02] Generación automática de una lista de la compra](https://github.com/GaelGoncalba/AutoShopping/milestone/2)
### Descripción: Desarrollar un módulo que utilice los datos de alimentos extraídos de las fuentes seleccionadas para generar automáticamente una lista de la compra que contenga todos los alimentos necesarios para satisfacer las necesidades nutricionales de una persona adulta durante 7 días, asegurando que la dieta sea lo más equilibrada posible.

### Entregables: 
- Módulo funcional que genera automáticamente una lista de la compra a partir de los datos extraídos.
- Documentación mínima en el propio código que describa los criterios de balance nutricional empleados y cómo se aplican en la generación de la lista.

### Producto: 
- Código del módulo base de la aplicación que genera una lista de la compra sin restricciones adicionales (en concreto, ajustes según presupuesto o preferencias dietéticas) de acuerdo con los valores recomendados para una dieta equilibrada.

### Viabilidad: 
  Este milestone es fundamental para validar que el módulo de generación de listas funciona correctamente y es capaz de producir resultados válidos. Será necesario:
- Verificar que las listas de compra cumplen con los estándares nutricionales básicos para una dieta saludable y equilibrada.
- Asegurar que la lógica del módulo se puede adaptar fácilmente a futuras mejoras y restricciones.

### Criterios de viabilidad: 
- La lógica de selección debe ser evaluada mediante pruebas nutricionales básicas, asegurando que los resultados cumplen con los mínimos recomendados para una dieta saludable (por ejemplo, equilibrio de macronutrientes).
- El código debe facilitar la posterior implementación de las funciones de _ajuste de lista de la compra a un presupuesto_ y _filtrado de los productos en listas de la compra_.

### HU relacionadas: [HU02]
----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M03] Ajuste de la lista de la compra a un presupuesto](https://github.com/GaelGoncalba/AutoShopping/milestone/3)
### Descripción: Implementar un módulo que, determinada por el usuario una cantidad de dinero que funciona como presupuesto para la compra semanal, utilice el algoritmo diseñado previamente para generar una lista de la compra que no supere en coste dicha cantidad.

### Entregables: 
- Módulo funcional que genera una lista de la compra ajustada a un presupuesto semanal definido por el usuario.
- Código que integre el presupuesto dentro de la lógica del algoritmo de selección de alimentos, priorizando los productos más accesibles sin comprometer el equilibrio nutricional.

### Producto: 
- Funcionalidad en la aplicación que permite al usuario ingresar un presupuesto semanal para la compra de alimentos, adaptando el código previamente generado para implementar la generación de listas de la compra sin restricciones de ningún tipo y aplicándole un algoritmo para evitar que el costo total de las listas generadas supere el presupuesto previamente determinado.

### Viabilidad: 
  Este milestone es fundamental para asegurar que la aplicación puede ajustarse a restricciones económicas sin comprometer la calidad nutricional de las listas de compra. Se deberá asegurar que:
- El módulo es capaz de generar listas dentro de los límites presupuestarios establecidos por el usuario.
- La lógica del algoritmo prioriza alimentos más económicos cuando es necesario, y sigue manteniendo el equilibrio nutricional.

### Criterios de viabilidad: 
- Las listas generadas deben respetar el presupuesto máximo ingresado por el usuario en el 100% de los casos.
- Se debe probar que el módulo puede generar múltiples combinaciones de productos que se ajusten al presupuesto sin comprometer los criterios del equilibrio nutricional, los cuáles fueron documentados en el código perteneciente a la funcionalidad de _generacion automatica de listas de la compra_.

### HU relacionadas: [HU01]
----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M04] Filtrado de las listas de la compra](https://github.com/GaelGoncalba/AutoShopping/milestone/4)
### Descripción: Implementar un módulo que cape la generación de listas de la compra según preferencias indicadas por el usuario, filtrando los productos disponibles por supermercados y tipo de alimento (pescado, verdura, carne...), y dando la posibilidad de no tener en cuenta algunos productos si pertenecen a alguno de esos filtros a la hora de incluirlos en dichas listas.

### Entregables: 
- Módulo funcional que filtra los productos según las preferencias del usuario, tales como el tipo de alimento o el supermercado donde se adquieren.
- Código que integre las opciones de exclusión de productos en el proceso de generación de la lista de la compra, permitiendo al usuario personalizar su lista según sus necesidades.

### Producto: 
- Funcionalidad que permite al usuario configurar y aplicar filtros antes de generar la lista de la compra, actualizando el módulo de generación de listas para que respete estos filtros y produzca una lista ajustada a las preferencias del usuario, y también implementando para ello nuevos apartados en la interfaz que permitan al usuario aplicar esos filtros de manera sencilla e intuitiva.

### Viabilidad: 
  Este milestone es esencial para personalizar la experiencia de usuario. Se deberá asegurar que:
- Los filtros aplicados por el usuario funcionan correctamente y se reflejan en la lista final generada.

### Criterios de viabilidad: 
- Las listas de compra generadas deben respetar al 100% las preferencias del usuario en cuanto a tipos de alimentos y supermercados seleccionados.
- Se debe probar que el módulo puede generar múltiples combinaciones de productos aplicando distintos filtros sin comprometer los criterios del equilibrio nutricional, los cuáles fueron documentados en el código perteneciente a la funcionalidad de _generacion automatica de listas de la compra_.

### HU relacionadas: [HU03]
