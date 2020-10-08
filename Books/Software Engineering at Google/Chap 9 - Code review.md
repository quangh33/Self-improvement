#### Benefits
- Check code correctness
- Ensure code change is comprehensible to other engineer
- Enforces consistency
- Promote team ownership
- Enable knowledge sharing
#### Process: 2 steps
- gaining an Look Good To Me (LGTM) from a peer engineer: check the correctness
- gaining approval from appropriate code readability reviewer: whether this change
is appropriate for their part of codebase
    - Will this code be easy to maintain?
    - Does it add to my technical debt?
    - Do we have the expertise to maintain it within our team?

> LGTM = the code does what it says + it is understandable


#### Best practices
1. Be polite and professional
- Trust - Respect - Humility
- expect prompt feedback (<24h)

2. Write small changes
- limit 200 LoC
- focus: e.g. the bug fix should focus solely on fixing the indicated bug
and updating the UT.

3. Write good change description
- What was changed?
- Why?
- Add comment if necessary

4. Keep reviewer to a Minimum
- 1 reviewer
- The cost of additional reviewers quickly outweighs their value.

5. Automate where possible
- Allow reviewers to focus on more important concerns than formatting

