FROM nginx:1.25-alpine

# Копируем кастомную HTML страницу
COPY index.html /usr/share/nginx/html/index.html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

