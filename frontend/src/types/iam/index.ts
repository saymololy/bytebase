import { readFileSync } from "fs";
import { load } from "js-yaml";
import { join } from "path";
import type { Permission } from "./permission";

const permissionYaml = readFileSync(join(__dirname, "permission.yaml"), "utf8");
const PERMISSION_DATA = load(permissionYaml) as { permissions: Permission[] };

export const PERMISSIONS: Permission[] = PERMISSION_DATA.permissions;

export * from "./permission";

export * from "./role";
