# Infrastructure

Local development is one Docker container: Postgres 16 on port **5433**. The API is a single Go process on port **8080**. There is no cache, event bus, API gateway, or cloud account in this repo yet.

## Why this is smaller than the hyperlocal project

That project runs PostGIS, Redis, Kafka, and Centrifugo locally, and its Terraform covers a VPC, EKS, RDS, ElastiCache, MSK, APISIX, and Centrifugo. Those pieces earn their place once several services share events and a live socket. A new product does not need them on day one.

## When something here is no longer enough

| Need | Add |
|------|-----|
| Shared environment | Amazon RDS for PostgreSQL, same schema as `backend/internal/platform/postgres/migrations` |
| Run the API somewhere other than a laptop | One container service (a small ECS service or a single Kubernetes deployment) |
| A second service that must react to the same change | An outbox table in Postgres first. Add a broker only when that table is not enough |
| Live server-to-client push | A push service such as Centrifugo, in front of one API, not a gateway mesh |
| Cache or geospatial queries | Redis, including ElastiCache if it is already on AWS |

Put new Terraform under `infrastructure/` when the first shared environment exists. Do not copy the hyperlocal module set ahead of that.
