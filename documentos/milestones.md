# Milestones del proyecto

## [[M01] Fundamentos del problema](https://github.com/GaelGoncalba/AutoShopping/milestone/1)
### Descripción: Determinar las tecnologías y lenguajes que se emplearán para extraer información de productos alimenticios desde la página web especificada y desarrollar un sistema básico que genere listas de la compra equilibradas de forma automática.

### Entregables: 
- Informe detallado que justifique las tecnologías elegidas (lenguajes de programación, frameworks, herramientas de scraping). 
- Código funcional del núcleo de la aplicación que implemente la extracción básica de datos y el análisis preliminar de listas de compra. 

### Producto: 
- Prototipo esquemático de la aplicación con una interfaz mínima que permita mostrar cómo se procesarían los datos de productos alimenticios y generar una lista de compra automática basada en criterios nutricionales.
- Pruebas iniciales con datos reales obtenidos de la web, que permitan analizar el potencial del algoritmo de balanceo de compras y evaluar la eficacia del scraping.

### Viabilidad: 
  Este milestone es crucial para establecer las bases tecnológicas del proyecto. Se deberá asegurar que:
- Las tecnologías seleccionadas son compatibles con los objetivos y escalabilidad del proyecto.
- Las fuentes de datos seleccionadas proveen información suficiente y confiable para alimentar la lógica del sistema.
- Se pueden generar listas de compra iniciales con los datos obtenidos.

### Criterios de viabilidad: 
- El prototipo debe demostrar que se pueden extraer y manipular datos de la página web objetivo.
- Las tecnologías deben ser evaluadas mediante pruebas de rendimiento y compatibilidad con los requisitos del proyecto, documentando las ventajas y desventajas de cada una. 
- El sistema debe generar una lista de compra equilibrada para al menos un conjunto de datos de prueba, demostrando que el enfoque es válido.

### HU relacionadas: [HU01],[HU02],[HU03]
----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M02] Generación automática de una lista de la compra](https://github.com/GaelGoncalba/AutoShopping/milestone/2)
### Descripción: Desarrollar un módulo que utilice los datos de alimentos extraídos de las fuentes seleccionadas para generar automáticamente una lista de la compra que contenga todos los alimentos necesarios para satisfacer las necesidades nutricionales de una persona adulta durante 7 días, asegurando que la dieta sea lo más equilibrada posible.

### Entregables: 
- Módulo funcional que genera automáticamente una lista de la compra a partir de los datos extraídos.
- Código documentado que explique la lógica usada para seleccionar los alimentos basados en los requisitos nutricionales generales.
- Reporte que describa los criterios de balance nutricional empleados y cómo se aplican en la generación de la lista.

### Producto: 
- Módulo base de la aplicación que genera una lista de la compra sin restricciones adicionales (por ejemplo, sin preferencias personales o limitaciones de presupuesto).
- Primer prototipo funcional de la lógica que evalúa los datos nutricionales y selecciona alimentos de acuerdo con los valores recomendados para una dieta equilibrada.

### Viabilidad: 
  Este milestone es fundamental para validar que el módulo de generación de listas funciona correctamente y es capaz de producir resultados válidos. Será necesario:
- Comprobar que los datos extraídos de las fuentes alimenticias son correctos y completos.
- Verificar que las listas de compra cumplen con los estándares nutricionales básicos para una dieta saludable y equilibrada.
- Asegurar que la lógica del módulo se puede adaptar fácilmente a futuras mejoras y restricciones (por ejemplo, ajustes según presupuesto o preferencias dietéticas).

### Criterios de viabilidad: 
- La lógica de selección debe ser evaluada mediante pruebas nutricionales básicas, asegurando que los resultados cumplen con los mínimos recomendados para una dieta saludable (por ejemplo, equilibrio de macronutrientes).

### HU relacionadas: [HU02]
----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M03] Ajuste de la lista de la compra a un presupuesto](https://github.com/GaelGoncalba/AutoShopping/milestone/3)
### Descripción: Implementar un módulo que, determinada por el usuario una cantidad de dinero que funciona como presupuesto para la compra semanal, utilice el algoritmo diseñado previamente para generar una lista de la compra que no supere en coste dicha cantidad.

### Entregables: 
- Módulo funcional que genera una lista de la compra ajustada a un presupuesto semanal definido por el usuario.
- Código que integre el presupuesto dentro de la lógica del algoritmo de selección de alimentos, priorizando los productos más accesibles sin comprometer el equilibrio nutricional.
- Documentación del algoritmo con ejemplos de listas generadas para diferentes rangos de presupuesto.

### Producto: 
- Funcionalidad en la aplicación que permite al usuario ingresar un presupuesto semanal para la compra de alimentos.
- Generación de una lista que respete tanto los requerimientos nutricionales como los límites monetarios, tomando en cuenta los precios de los alimentos obtenidos de las fuentes de información.

### Viabilidad: 
  Este milestone es fundamental para asegurar que la aplicación puede ajustarse a restricciones económicas sin comprometer la calidad nutricional de las listas de compra. Se deberá asegurar que:
- El módulo es capaz de generar listas dentro de los límites presupuestarios establecidos por el usuario.
- La lógica del algoritmo prioriza alimentos más económicos manteniendo el equilibrio nutricional.

### Criterios de viabilidad: 
- Las listas generadas deben respetar el presupuesto máximo ingresado por el usuario en el 100% de los casos.
- Se debe probar que el módulo puede generar múltiples combinaciones de productos que se ajusten al presupuesto sin comprometer el equilibrio nutricional.

### HU relacionadas: [HU01]
----------------------------------------------------------------------------------------------------------------------------------------------------------------
## [[M04] Filtrado de las listas de la compra](https://github.com/GaelGoncalba/AutoShopping/milestone/4)
### Descripción: Implementar un módulo que cape la generación de listas de la compra según preferencias indicadas por el usuario, filtrando los productos disponibles por supermercados y tipo de alimento (pescado, verdura, carne...), y dando la posibilidad de no tener en cuenta algunos productos si pertenecen a alguno de esos filtros a la hora de incluirlos en dichas listas.

### Entregables: 
- Módulo funcional que filtra los productos según las preferencias del usuario, tales como el tipo de alimento o el supermercado donde se adquieren.
- Código que integre las opciones de exclusión de productos en el proceso de generación de la lista de la compra, permitiendo al usuario personalizar su lista según sus necesidades.

### Producto: 
- Funcionalidad que permite al usuario configurar y aplicar filtros antes de generar la lista de la compra.
- Actualización del módulo de generación de listas para que respete estos filtros y produzca una lista ajustada a las preferencias del usuario.

### Viabilidad: 
  Este milestone es esencial para personalizar la experiencia de usuario. Se deberá asegurar que:
- Los filtros aplicados por el usuario funcionan correctamente y se reflejan en la lista final generada.
- Los filtros son lo suficientemente amplios y prácticos para cubrir las necesidades comunes de los usuarios.

### Criterios de viabilidad: 
- Las listas de compra generadas deben respetar al 100% las preferencias del usuario en cuanto a tipos de alimentos y supermercados seleccionados.
- Los filtros deben ser probados con diferentes combinaciones de criterios, asegurando que excluyen correctamente los productos no deseados.

### HU relacionadas: [HU03]
