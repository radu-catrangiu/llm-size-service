"use strict";

const { huggingface: hf } = require("./config");
const { execSync } = require("node:child_process")
const { existsSync, rmSync } = require("node:fs")
const path = require("node:path")
const { randomUUID } = require("node:crypto")

const tmpDir = "/tmp"

async function handler(model) {
    const repoDirName = `repo-${randomUUID()}`;
    const repository = `https://:${hf.password}@huggingface.co/${model}`
    const destPath = path.join(tmpDir, repoDirName);
    if (existsSync(destPath)) {
        rmSync(destPath, { force: true, recursive: true });
    }

    const cloneCmd = `git clone --no-checkout ${repository} ${repoDirName}`
    execSync(cloneCmd, {
        cwd: tmpDir,
    });

    const cmd = `git-lfs ls-files -s --json`
    const resultStr = execSync(cmd, { cwd: destPath });
    
    const result = JSON.parse(resultStr);

    rmSync(destPath, { force: true, recursive: true });

    const sizeMB = result.files.reduce((acc, cur) => {
        return acc + (cur.size / (1000 * 1000))
    }, 0);

    const roundedGi = Math.ceil(sizeMB / 1000);

    return {
        size: `${sizeMB / 1000}Gi`,
        roundedSize: `${roundedGi}Gi`,
    };
}

module.exports = handler;
