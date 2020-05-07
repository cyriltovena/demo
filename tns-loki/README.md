TURN OFF NOTIFICATIONS.

Show of hands who is using grafana ?
2 Part demo Grafana and LogCLI
You can use the link to download all the configuration that I've used to setup do this demo


I've already deployed Loki using helm. I actually use the loki-stack chart which an umbrella chart that can allow you to deploy loki promtail and prometheus already configured and ready to go. It will automatically add datasource. It's basically one command to install.

For this demo we're going to use a demo application that's we built.
It's classic 3 tiered application with a load balancer/proxy, the app itself and database. Very common scenario.

helm install --name loki-stack loki/loki-stack -f values.yaml --namespace monitoring
kubectl get pods -n monitoring

Now if you don't run on kubernetes no problem, Loki like we said before can run a single binary. You just need to download a release and start it with a configuration file.

curl -O -L  https://github.com/grafana/loki/releases/download/v1.4.1/loki-darwin-amd64.zip

curl http://localhost:3100/ready

You can also run it within docker too, with a simple docker run.

Present the APP http://localhost:30030/ We received an alert and this app is having 500.

Show datasource already installed.

Go to the dashboard explain the problem, then explore metrics and switch to Loki
Use a filter to reduce the noise

Live stream errors

Panic Detail Panel and Show context.

Show a count_over_time & rate

   count_over_time({namespace="tns"} |= "level=error"[5m])
    /
    count_over_time({namespace="tns"}[5m])
    * 100

Build a dashboard with The rate of errors by applications and the logpanel

Demo logcli :

Why would you use the logcli ?

logcli query '{ namespace="tempo-dev", name="ingester" } |= "level=error"' -o raw --since=24h --limit=5000 | sed -e 's/[0-9]//g' | sort | uniq -c

logcli query '{ job="default/eventrouter" } |= "OOM"' -o raw |  jq -r '(.event.reason) + "\t" + (.event.involvedObject.namespace) + "/" + (.event.involvedObject.name) + "\t" + (.event.message)'

Conclude
