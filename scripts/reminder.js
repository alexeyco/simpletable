module.exports = async function runReminder({ github, context, core }) {
  const title = "Refresh the README comparison";
  const { owner, repo } = context.repo;
  const openIssues = await github.paginate(github.rest.issues.listForRepo, {
    owner,
    repo,
    state: "open",
    per_page: 100,
  });
  const reminderExists = openIssues.some(
    (issue) => !issue.pull_request && issue.title === title,
  );

  if (reminderExists) {
    core.info(`An open issue already exists: ${title}`);
    return;
  }

  await github.rest.issues.create({
    owner,
    repo,
    title,
    body: "Review the public evidence in the README comparison, update any claims that changed, and refresh the last-checked date.",
  });
};
