# Discovery

## Questions to answer

1. Which primitives are we checking?
    - Only the functions for setting max open / idle connections, max idle timeout?
    - Or do we check anything else?
2. What if the value comes from config? Can a static checker resolve to what it is set?
    - E.g., if the value comes from env variable in Kubernetes workloads.
3. What if a project contains multiple applications and only one of them uses a database?
    - How does the linter know which one to check? User to pass in a target for starters?
    - Similarly, if there are multiple applications and they both uses a database
        - How do we ensure all of them adhere to the requirement?
        - What if the requirements differ between apps?
        - What if they use different databases?
4. There are a lot of ORM and libraries alike to communicate with a database.
    - Which one are we going to be looking for?
5. There are many databases, which one are we doing this for?
6. Can we be slow? Or do we need to be fast?

## Assumptions and limitations

1. Assume we are checking for function calls to set max open / idle conns and max idle timeout.
2. Assume parameter in functions we will be performing checks for is set in in-line code, and not passed in from config.
3. Assume one application and one database type for starters.
4. Assume we are doing this for PostgreSQL database.
5. Assume sql package for starters
6. Be slow as a start

## Requirements

Start with:

- Detect sql.SetMaxOpenConns;
- Detect the number used;
- Check that it matches a user provided value;
- Fail if that value doesn't match;
- Fail if sql.SetMaxOpenConns isn't called.
