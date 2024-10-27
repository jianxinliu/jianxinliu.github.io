const esbuild = require("esbuild");
const fs = require("fs");

const branch = process.env.BRANCH;
const isProd = process.env.NODE_ENV === "PROD";
console.log("env: ", process.env.NODE_ENV);
console.log("branch: ", branch);

const version = process.env.VERSION || "init";

const outFilePath = isProd
  ? `./build/bundle-${version}${branch ? "-" + branch : ""}.js`
  : `./build/bundle_dev.js`;

esbuild.buildSync({
  entryPoints: ["./src/index.ts"],
  bundle: true,
  platform: "browser",
  outfile: outFilePath,
  minify: isProd,
  tsconfig: "./tsconfig.json",
  dropLabels: isProd ? ["DEV"] : [],
  drop: isProd ? ["console", "debugger"] : [],
  treeShaking: true,
  define: {
    // 注入版本号
    "process.env.VERSION": `"${version}"`,
  },
});

// remove tfjs boundle comment
const content = fs.readFileSync(outFilePath).toString();

fs.writeFileSync(
  outFilePath,
  content.split("/*! Bundled license information:")[0]
);

if (isProd) {
  // 方便 sdk 通过 function() { js } 的方式进行注入
  const content = fs
    .readFileSync(outFilePath)
    .toString()
    .replace(/^"use strict";\(\(\)=>\{/, "")
    .replace(/\}\)\(\);\n?$/, "\n");

  fs.writeFileSync(outFilePath, content);
}
