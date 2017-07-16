# E-caiu?

[PT-br]

O sistema *E Caiu?* verificará a disponibilidade dos sites da UFPB. No arquivo de log.txt é disponibilizado algumas informações úteis, por exemplo, a data da verficação juntamente com a hora e com isso poderá medir por quanto a tempo o determinado site ficou fora do ar.

## funcionalidades

- [x] Verifica se os sites que estão salvo no arquivo sites.txt estão no ar
- [x] salva em arquivo os logs da aplicação
- [ ] web api - em andamento - [@abraaohonorio](https://github.com/AbraaoHonorio/)
- [ ] deploy da aplicação em andamento - [@abraaohonorio](https://github.com/AbraaoHonorio/)
- [ ] Integração com Banco de dados
- [ ] Front end do sistema
- [ ] Integração com sistema troca de mensagem(slack, email, messenger...)
- [ ] Analise dos dados
  - [ ] Indentificar qual foi o site ficou fora ar.
  - [ ] Quanto tempo o sistema ficou fora do ar
  - [ ] Em que dia da semana  
  - [ ] porcetagem que de quanto tempo o sistema ficou disponvel, relação diaria,semanal,mensal

### Building no projeto
  ```sh
    go run main.go
  ```
  
## Contribuições são bem-vindas!

No tópico de [funcionalidades](https://github.com/AbraaoHonorio/E-caiu/#funcionalidades) criei as features que irei implementar para o projeto, caso tenha alguma feature interessante é só da um pull request no próprio README, caso queira implementar a feature, por favor colocar a tag em andamento - seu github ( em andamento - @AbraaoHonorio)

###### Exemplo de log
       "16/07/2017 15:00:18 - https://sigaa.ufpb.br/sigaa/verTelaLogin.do -Online: true"
