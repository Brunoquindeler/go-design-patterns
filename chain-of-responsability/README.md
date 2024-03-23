# Chain of Responsibility em Go

[Refactioring Guru](https://refactoring.guru/pt-br/design-patterns/chain-of-responsibility/go/example)

## O **Chain of Responsibility** é um padrão de projeto comportamental que permite passar a solicitação ao longo da cadeia de handlers em potencial até que um deles lide com a solicitação.

O padrão permite que vários objetos tratem a solicitação sem acoplar a classe remetente às classes concretas dos destinatários. A cadeia pode ser composta dinamicamente em tempo de execução com qualquer handler que siga uma interface de handler padrão.

## Exemplo conceitual

Vejamos o padrão Chain of Responsibility com o caso de um aplicativo de hospital. Um hospital pode ter vários departamentos, como:

* Recepção
* Médico
* Sala de remédios
* Caixa

Sempre que chega algum paciente, ele vai primeiro para a Recepção, depois para o Médico, depois para a Sala de Remédios e depois para o Caixa (e assim por diante). O paciente está sendo enviado por uma cadeia de departamentos, onde cada departamento o envia mais adiante na cadeia, uma vez que sua função esteja concluída.

O padrão é aplicável quando há vários candidatos para processar a mesma solicitação. Quando você não quer que o cliente escolha o receptor, pois vários objetos podem lidar com a solicitação. Além disso, você deseja desacoplar o cliente dos receptores. O cliente só precisa conhecer o primeiro elemento da cadeia.

Como no exemplo do hospital, o paciente primeiro vai até a recepção. Então, com base no status atual do paciente, a recepção envia para o próximo handler da cadeia.