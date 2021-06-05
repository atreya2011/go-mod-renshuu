# go-mod-renshuu

<https://github.com/google-github-actions/release-please-action/issues/78#issuecomment-708836986>

Solved "No merged release PR found" with above solution

1. `git checkout -b release-v0.0.1`
2. `git commit --allow-empty -m "chore: release 0.0.1"`
3. Create and merge PR
4. Subsequent merged PRs will create the release PR
5. Profit!
