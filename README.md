# Time management

Eisenhower's Urgent/Important Principle\
Using Time Effectively, Not Just Efficiently

## Schema
```sql
-- Table Definition
CREATE TABLE "public"."task" (
    "id" varchar NOT NULL,
    "title" varchar NOT NULL,
    "description" text NOT NULL,
    "is_done" bool DEFAULT false,
    "is_over" bool DEFAULT false,
    "matrix" varchar NOT NULL,
    "duration" int4 NOT NULL,
    "start_at" int4 NOT NULL,
    "due" int4 NOT NULL,
    "created_at" int4 NOT NULL,
    "updated_at" int4 NOT NULL,
    PRIMARY KEY ("id")
);
```

## Examples request (endpoints)
https://github.com/SemmiDev/Nugaz/blob/main/request.http
