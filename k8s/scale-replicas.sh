CURRENT=$(kubectl get deployment argocd-application-controller \
  -n argocd \
  -o jsonpath='{.spec.replicas}')
echo "Current replica count is $CURRENT"

# Pause
kubectl scale deployment argocd-application-controller \
  -n argocd \
  --replicas=0

# …when ready to resume…
kubectl scale deployment argocd-application-controller \
  -n argocd \
  --replicas="$CURRENT"
