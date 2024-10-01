# AWS CRON

At AWS a cronjob can be implemented in many ways, but probably, the most popular way is to trigger a job (lambda or Fargate task) by creating an EventBridge rule with a schedule.

So this cli command can list all such scheduled jobs.

## Example of command output

```bash
┌───────────────────┬──────────┬───────────────────────────────────────────────────────────┐
│      Schedule     │  Status  │                             Name                          │
├───────────────────┼──────────┼───────────────────────────────────────────────────────────┤
│ cron(0 7 * * ? *) │ ENABLED  │ automation-test-seed-data                                 │
│                   │          │ Fires once a day to trigger the automation-test-seed-data │
│                   │          │ job                                                       │
├───────────────────┼──────────┼───────────────────────────────────────────────────────────┤
│ cron(0 8 * * ? *) │ ENABLED  │ automation-test                                           │
│                   │          │ Fires once a day to trigger the automation-test job       │
├───────────────────┼──────────┼───────────────────────────────────────────────────────────┤
│ rate(5 minutes)   │ ENABLED  │ messages-fanout                                           │
│                   │          │ Triggers job to deliver pending messages                  │
├───────────────────┼──────────┼───────────────────────────────────────────────────────────┤
│ cron(0 1 * * ? *) │ DISABLED │ send-report                                               │
│                   │          │ Fires once a day to trigger job that send daily report    │
└───────────────────┴──────────┴───────────────────────────────────────────────────────────┘
```

## Usage

```bash
aws-cron --help

# show list of cron jobs
aws-cron
aws-cron --output=table
aws-cron -o=text

# show cli version
aws-cron version
aws-cron version --output=json

# todo: trigger cronjob
aws-cron <name/arn?> trigger
```

## Release

```bash
git tag -a v0.1.3 -m "Release v0.1.3"
git push origin v0.1.3

```