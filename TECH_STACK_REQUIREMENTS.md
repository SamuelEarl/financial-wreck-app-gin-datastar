# Requirements List for Financial Wreck Tech Stack

• Low maintenance and exceptional error reporting during development - not only at build time. 
    • Statically typed language or extensive use of static analyzers. 
    • Display build errors during development.
• SSR capable for SEO, but also must provide some dynamic client-side capabilities. I want to use this for both dynamic web apps and e-commerce sites.
• I have to enjoy the programming language and framework. 
• A Python backend would be cool for AI and ML capabilities. This is not required, though, because I can call a Python service, if necessary.
• I would like the server to be fast so it can handle many requests without too much horizontal scaling. 
• I like opinionated languages and framework (like Python frameworks), but I didn't like them to be so opinionated that they make the entire tech stack decision for you (Rails, Phoenix). I have some opinions and preferences too and I don't always like what the core framework developers like.
• I want the apps to be server full instead of serverless. Apparently serverless apps can end up being more expensive than server full. With many of the meta frameworks the server seems to be an afterthought. That means that the middleware is difficult to use (which makes authentication harder to implement), databases are difficult to connect to unless they are serverless and have a serverless driver, and other features seem to be missing (e.g. OpenAPI documentation).
• A tech stack that does not need to be updated constantly. I hate how the JavaScript ecosystem moves so rapidly. That is a big reason for JavaScript fatigue and burnout. I want my tech stack to be "evergreen" as much as possible so I don't have to update or change any part of the stack very often. I want to spend my time adding new features to my apps and being productive rather than going back and updating parts of the stack to the latest versions and making sure that everything is still working together properly after the updates, which is what you often have to do with a tech stack that is JavaScript-heavy.

## I am leaning toward the following tech stack

• Server: Go (Gin framework)
• Client: Templ templating engine, Datastar, Lightning CSS, DatastarUI (use as an example of how to create components with Datastar and Templ using AI chat bots)
• NeonDB or FalkorDB
