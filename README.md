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
- Save Button, Open Button, New Button.

## Step 4:

Creating the pdf export, a huge yackshave, but lets gooo!

- Pdf renderer has a really interesting and useful abstraction with the Walk function and the abillity to skip children.
- It's not feature complete, but we have a good base of links, headings etc etc.
- Wiring up the export function is pretty straightforward!

It really ain't pretty yet, lets make something cool...


## Step 5:

Eat your vegetables, time to setup CI/CD and make a release! Yea, you should probably do this as step 1, but being true to the reality of how i worked on this, I'm putting it as step 5.

Here is the high level steps I took a look through..

- Figuring out the production build, windows/linux cross build and copying the binary for later
- Figuring out the release process. https://github.com/softprops/action-gh-release 
- Push your code about six times to get build syntax right
- attempt to use a provided github action to matrix build, works for nix but not for windows due to obscure reasons. I think for this simple app a cross compile should be fine, so I'm just going to go with a cross compiled app for simplicity