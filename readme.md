Crawler feito em Golang

o que e um Crawler ? 

    É um software que se encarrega de percorrer a todos os links das páginas webs de forma sistemática e automática. Para definir a função do Crawler, podemos dizer que ele rastreia e analisa toda o site, seguindo por todos os links internos e externos.

Como usar?

    Necessario o uso do Docker conectado ao mongodb

como criar a base de dados? Apos pegar o codigo digite os seguintes comandos no terminal:

    docker run -d --name mongodb -p 27017:27017 mongo (para criar o container)
    docker start mongodb (para iniciar)
    docker ps (container info)

para acessar o banco de dados:

    docker exec -it mongodb /bin/bash
    mongo
    show dbs;
    use crawler
    show collections;

para filtrar: 

    db.links.count({})  (mostra quantos sites foram visitados)
    db.links.find({})   (mostra os sites visitados com logs)

Para usar o Crawler:

    go run main.go --url=https://github.com (URL QUE DESEJA PESQUISAR)
    
    
https://user-images.githubusercontent.com/86691253/175209408-3860f206-429d-4bcd-bf70-43583662fabe.mp4


