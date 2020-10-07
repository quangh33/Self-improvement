#### Question docs answer
- Why were these design decision made?
- Why did we implement this code in this manner?
- Why did I implement this code in this manner, if you're looking at your own code 2 years later?

Benefit of docs is not immediate => it is often considered as a burden
#### Docs should:
- have internal policies to be followed
- be place under source control
- clear ownership
- undergo review for changes
- have issues tracked
- periodically evaluated

#### Docs types
- ref docs, including code comment
    - file comment: what is contained in the code you are reading
    ```cpp
    // -----------------------------------------------------------------------------
    // str_cat.h
    // -----------------------------------------------------------------------------
    //
    // This header file contains functions for efficiently concatenating and appending
    // strings: StrCat() and StrAppend(). Most of the work within these routines is
    // actually handled through use of a special AlphaNum type, which was designed
    // to be used as a parameter type that efficiently manages conversion to
    // strings and avoids copies in the above operations.
    ```
    > Any API that cannot be succinctly described in the first paragraph or two is usually the sign of an API that is not well thought out
    - class comment: describing the class/struct, important methods of that class, and the purpose of the class
    ```cpp
    // -----------------------------------------------------------------------------
    // AlphaNum
    // -----------------------------------------------------------------------------
    //
    // The AlphaNum class acts as the main parameter type for StrCat() and
    // StrAppend(), providing efficient conversion of numeric, boolean, and
    // hexadecimal values (through the Hex type) into strings.
    ```
    - function comment: describing what the function does
    
- design docs
    - required approved design doc before starting any major project
    - Google docs for collaboration
    - should cover:
        - design goals
        - alternative designs and trade-off
- tutorials
    - the best time to write a tutorial, if one does not yet exist, is when you first join a team.
    - separate input/output in a separated line using **`monospace bold`**
    - combine all atomic user operations into single steps
- conceptual docs
    - deeper explanations or insights that cannot be obtained by reading ref docss
    - may lead to duplication but with purpose: **promote clarity**
    - needs to be useful to a broad audience: both experts and novices alike. 

- landing pages
    - team page
    - only purpose: navigate    
    > If something on a landing page is doing more than being a traffic cop, it is not doing its job
    - separate: `goto` page for customer (of your API) and `team page` (for internal usage)                                                          
                                                                
> Some documentation styles (and some documentation generators) require various forms of boilerplate on function comments, like “Returns:”, “Throws:”, and so forth, but at Google we haven’t found them to be necessary. It is often clearer to present such information in a single prose comment 
>that’s not broken up into artificial section boundaries:

```cpp
// Creates a new record for a customer with the given name and address,
// and returns the record ID, or throws `DuplicateEntryError` if a
// record with that name already exists.
int AddCustomer(string name, string address);
```

#### Docs review
- technical review: for accuracy
- audience review: for clarity
- writing review: for consistency


#### Docs philosophy
1. WHO, WHAT, WHEN, WHERE, and WHY
=> try to answer those questions in first 2 paragraphs 
- Who: audience
> This document is for new engineers on the Secret Wizard project.

- What: purpose
> This document is a tutorial designed to start a Frobber server in a test environment.
- When: when this document was created, reviewed, or updated.
- Where:  often implicit as well, but decide where the document should live.
- Why:
    - Summarize what you expect someone to take away from the document after reading it
    - in the introduction to a document

2. The Beginning, Middle, and End
- in documentation, redundancy is often useful
- docs should often have, at a minimum, those three sections.

3. Parameters of Good Doc
- completeness
- accuracy
- clarity
- good doc == doing its intended job
