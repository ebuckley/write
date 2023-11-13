# A writer application

This is created as a step by step guide in how to build using the wails framework.

I'm choosing my favorite SPA fotm (framework of the month), but you could do this with anything else.

IF you're interested in using tailwind and react to make a tiny bundle for a frontend app, read on!

## Step 1: Create a new project

```bash
wails init -n write -t https://github.com/hotafrika/wails-vite-react-ts-tailwind-template
cd writer
```


## Step 2:

- "Design" the application: https://v0.dev/t/OpQCijO6Hm8
- Implement the markdown renderer, see lib/markdown.go and `app.SendMarkdownToRenderer`
- Now you can write markdown!

## Step 3:

- Implement all the other buttons etc etc
- 