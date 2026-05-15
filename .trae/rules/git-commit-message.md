---
alwaysApply: true
scene: git_message
---

### **Git Commit Message Style Guide (Concise Template)**

All commit messages must be in **pure English**, with no Chinese characters.

#### **Format**
```
<TYPE>(<scope>): <subject>

<body>

<footer>
```

#### **Rules**
*   **Header**: Mandatory. `<type>(<scope>): <subject>`
    *   `<type>`: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `chore`.
    *   `<scope>`: Optional. Module/file affected (e.g., `api`, `ui`).
    *   `<subject>`: Imperative, concise summary (< 50 chars). No period.
*   **Body**: Optional but encouraged for complex changes. Explain *what* and *why*. Wrap at 72 chars.
*   **Footer**: For `BREAKING CHANGE:` or issue references (e.g., `Closes #123`).
*   Separate header from body with a blank line.

#### **Examples**
```
feat(auth): add OAuth2 login support

fix(api): prevent null ptr in user fetch
```
---