# Requirements

Right now, dblinter MUST:

- Detect sql.SetMaxOpenConns;
- Detect the number used;
- Check that it matches a user provided value;
- Fail if that value doesn't match;
- Fail if sql.SetMaxOpenConns isn't called.
