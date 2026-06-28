# Contributing to Babank

First off, thank you for considering contributing to Babank! 

## Branching Strategy

We strictly follow a **Feature Branching** workflow. The `main` branch is our single source of truth and is highly protected. You should never commit directly to `main`.

### Workflow
1. **Create a Feature Branch:** Always branch off of the latest `main` branch. Name your branch descriptively based on the feature or bug you are working on.
   - Good: `feature/user-profile-ui`
   - Good: `bugfix/login-timeout`
   - Bad: `my-changes`

2. **Develop and Commit:** Make your changes locally. Ensure you include relevant tests and that your code passes all linting rules (`golangci-lint` or `angular-eslint`).

3. **Open a Pull Request:** Push your feature branch to the remote repository and open a Pull Request (PR) against `main`. 
   - You must fill out the Pull Request Template entirely.
   - You must include test coverage summaries.

4. **Code Review:** All PRs require at least one approving review before they can be merged. Address any feedback from reviewers.

5. **Merge:** Once your PR is approved and all CI status checks pass, it will be squash-merged into `main`. The feature branch will be automatically deleted post-merge.

## Development Setup
*(Instructions for setting up the local Bazel monorepo will be added here).*
