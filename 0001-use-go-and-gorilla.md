# ADR 0001: Use Go with Gorilla

**Date:** 2026-02-09  
**Status:** Accepted

## Context

The project is a small legacy application that needs to be rewritten in another language.  
The goal of the rewrite is to keep the system simple, understandable, and easy to operate, while documenting architectural decisions using ADRs.

The project scope is limited, and the group is still learning both the technology stack and the ADR process. A lightweight web solution is required to handle HTTP requests without introducing unnecessary complexity.

## Decision

We decided to implement the rewritten application in **Go**, using the Go standard library together with **Gorilla** for HTTP routing. 

## Reasoning

- Go provides a simple and explicit programming model with strong built-in support for HTTP servers and is widely used in the industry.
- Gorilla Mux improves routing and middleware support while remaining lightweight and close to the Go standard library.
- This approach avoids the complexity of full web frameworks and keeps the focus on the rewrite and DevOps practices.

## Consequences

- The application will be easy to deploy and operate due to Go’s single-binary output.
- Routing logic will be clearer and more maintainable than using the standard library alone.
