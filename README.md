# GitConnect – DevOps CI/CD Project  

## 📌 Project Overview  
This project demonstrates a full DevOps workflow by containerizing and deploying a MERN-based GitConnect application.  
It covers CI/CD pipelines, Kubernetes deployment, and monitoring using Prometheus & Grafana.  

## ⚡ Tech Stack  
- Frontend: React  
- Backend: Node.js / Express  
- Database: MongoDB  
- Containerization: Docker, Docker Compose  
- Orchestration: Kubernetes (Minikube / k3d)  
- CI/CD: GitHub Actions  
- Monitoring: Prometheus & Grafana  

## 🚀 Features  
- CI/CD pipeline with automated build & deployment  
- Kubernetes manifests for backend, frontend, services, and secrets  
- Monitoring stack (Prometheus + Grafana)  
- Scalable deployment with NodePort services  

## 🌍 Live Demo  
- Frontend: [http://your-frontend-link.com](http://your-frontend-link.com)  
- Backend API: [http://your-backend-link.com](http://your-backend-link.com)  
- Grafana Dashboard: [http://your-grafana-link.com](http://your-grafana-link.com)  
- Prometheus: [http://your-prometheus-link.com](http://your-prometheus-link.com)  

## 📊 Monitoring  
- Prometheus: Collects metrics  
- Grafana: Visualizes metrics & dashboards  

## 📷 Screenshots (To Add)  
- Frontend app running  
- Grafana dashboard with metrics  
- Architecture diagram  

## 📐 Architecture Diagram  
```text
   GitHub Actions (CI/CD)
             |
     DockerHub Registry
             |
      ------------------
      |                |
  Frontend (React)   Backend (Node)
      |                |
      ----- Kubernetes -----
             |  
       Prometheus + Grafana
