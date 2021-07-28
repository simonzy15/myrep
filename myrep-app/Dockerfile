# Stage 1: Compile and Build angular codebase

# Use official node image as the base image
FROM node:14.17-alpine as build

# Set the working directory
WORKDIR /usr/src/app

# Add the source code to app
COPY ./ /usr/src/app/

# Install all the dependencies
RUN npm install

# Generate the build of the application
RUN npm run build


# Stage 2: Serve app with nginx server

# Use official nginx image as the base image
FROM nginx:latest

# Copy the build output to replace the default nginx contents.
COPY --from=build /usr/src/app/dist/myrep-app /usr/share/nginx/html

# Expose port 4200
EXPOSE 4200