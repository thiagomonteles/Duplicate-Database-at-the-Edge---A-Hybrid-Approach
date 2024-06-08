# Project: Edge-Replicated Databases - A Hybrid Approach

## Students:
- Ben Hur Faria Reis
- Filipe Silveira Chaves
- Thiago Monteles de Sousa
- Felipe Aguiar Costa
- João Paulo Rocha

**Context:**
In the context of database replication, strong consistency or eventual consistency approaches offer different guarantees, where in some cases, one or the other is more recommended. In this sense, the project aims to create a replicated database so that both consistency approaches are implemented, and based on usage, the user can choose which one they prefer to use.

**Objectives:**

- Develop an edge-replicated database that offers both strong consistency and eventual consistency.
- Explore the use of techniques such as Conflict-free Replicated Data Types (CRDT) to ensure eventual consistency.
- Use the Raft algorithm to ensure strong consistency.
- Provide users with the ability to select the consistency approach that best suits their needs and the context of use of the replicated database.

## Project Stages:

### To complete the project, we will have the following stages:

1. Literature review: (ALL)
2. Preparation of internal presentations on consistency methods: (ALL)
3. Decision on consistency methods: (ALL)
4. Development of an architectural scheme of communication and data: (ALL)
5. Coding of a replicated database: (Felipe Aguiar, Ben Hur)
6. Coding of the CRDT protocol: (Felipe Aguiar)
7. Coding of the RAFT protocol: (Filipe Chaves, João Paulo)
8. Coding of a node state viewer: (Thiago, Ben Hur)
9. Coding of the final aggregator of protocols and replicated database
