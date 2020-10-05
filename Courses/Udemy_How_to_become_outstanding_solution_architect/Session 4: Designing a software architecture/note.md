1. S.A requirements
- user requirement: what users want
- business requirement: cheaper, faster ...?
- IT requirement: cloud or on-premise

2. Architecture
- elements
- how elements behave in collaboration
- composition of elements into larger subsystem
- architecture style that guides this composition: pattern

3. Goals:
- document high-level structure
- not go into detail: just black box
- **minimize complexity** <= break down the design into layers, SRP
- address all requirement
- compatible with all use case

4. Key principles
- Separation of concerns: layers
- SRP: each element has one responsibility
- Principle of least knowledge: elements have no/ minimum knowledge about other's detail <= use **Interface**
- DRY: no multiple elements with same responsibility
- Minimize upfront design: start small

5. Guideline
- Establish code convention (co-dev with Tech lead)
- **Composition over inheritance** (co-dev with Tech lead)
- Layers:
    - Separate areas of concern
    - Define communication b/w layers
- Component:
    - No component rely on internal of another: **use Interface**
    - SRP
    - Put system-wide component into system-wide layer 
