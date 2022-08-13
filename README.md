Run ```docker compose up```<br />
This run should end up in an error, as database has not been initiated<br />
Enter either the user, product, or transaction service container with ```docker exec -it <Container_ID> /bin/bash```<br />
Run ```make createdb``` to create db<br />
Run ```make migrate``` to migrate the tables<br />
Exit the container and stop the container re-run ```docker compose up```<br />
Run ```docker compose up``` again<br />
Please make sure that Port _5432_ is available for PgSQL and port _8000_, _8001_, _8002_, _8003_, _8004_ and _8005_ are available
