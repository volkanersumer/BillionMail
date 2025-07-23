import { execSync } from 'child_process';

async function getPull() {
    try {
        // 在当前目录执行git pull命令，并同步输出结果
        const output = execSync('git pull', { stdio: 'inherit', cwd: process.cwd() });
        console.log('git pull 执行成功');
        return true
    } catch (error) {
        console.error('git pull 执行失败:', error.message);
        process.exit(1);
        return false
    }
}

getPull()