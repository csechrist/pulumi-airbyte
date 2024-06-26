// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceSftpBulk DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceSftpbulk = airbyte.getSourceSftpBulk({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceSftpBulk(args: GetSourceSftpBulkArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceSftpBulkResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceSftpBulk:getSourceSftpBulk", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceSftpBulk.
 */
export interface GetSourceSftpBulkArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceSftpBulk.
 */
export interface GetSourceSftpBulkResult {
    /**
     * The values required to configure the source. Parsed as JSON.
     */
    readonly configuration: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly sourceId: string;
    readonly sourceType: string;
    readonly workspaceId: string;
}
/**
 * SourceSftpBulk DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceSftpbulk = airbyte.getSourceSftpBulk({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceSftpBulkOutput(args: GetSourceSftpBulkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceSftpBulkResult> {
    return pulumi.output(args).apply((a: any) => getSourceSftpBulk(a, opts))
}

/**
 * A collection of arguments for invoking getSourceSftpBulk.
 */
export interface GetSourceSftpBulkOutputArgs {
    sourceId: pulumi.Input<string>;
}
