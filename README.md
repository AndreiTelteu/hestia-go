
The intent behind this project is for me to get real-world experience with GoLang and SolidJS.

The scope of this project is to rebuild as much as I can from the [hestiacp/hestiacp](https://github.com/hestiacp/hestiacp/) project, while providing a faster, safer, powerful and customizeable experience.

Things I want to acomplish:

- build a structure that bundles all the frontend assets in a single go binary **✅ DONE**
- build a plugin structure that allows building an app store in the future **✅ DONE**
- rebuild the entire frontend of hestiacp using my [solidjs-window-manager](https://github.com/AndreiTelteu/solidjs-window-manager) plugin
- I don't know yet if I want to rewrite all the hestiacp shell scripts or use theirs and build this backend around it
- improove the nginx configuration cababilities
- make a plugin for supervisor
- make a live logs interface
- build a mobile app !

Demo:
```bash
docker compose up -d
# this request will be served from the frontend vite dev server
curl http://localhost:3000/
# this request will be served from core
curl http://localhost:3000/api
# this request will be served from the example-plugin via static filesystem
curl http://localhost:3000/api/example-plugin/
# this request will be served from the example-plugin custom api handlers
curl http://localhost:3000/api/example-plugin/plugin
```



