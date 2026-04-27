#!/bin/zsh


git add .
git commit -m "update"
git push origin main


git tag v0.4.1
git push origin v0.4.1