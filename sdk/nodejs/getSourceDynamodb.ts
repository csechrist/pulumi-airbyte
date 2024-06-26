// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceDynamodb DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceDynamodb = airbyte.getSourceDynamodb({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceDynamodb(args: GetSourceDynamodbArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceDynamodbResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceDynamodb:getSourceDynamodb", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceDynamodb.
 */
export interface GetSourceDynamodbArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceDynamodb.
 */
export interface GetSourceDynamodbResult {
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
 * SourceDynamodb DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceDynamodb = airbyte.getSourceDynamodb({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceDynamodbOutput(args: GetSourceDynamodbOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceDynamodbResult> {
    return pulumi.output(args).apply((a: any) => getSourceDynamodb(a, opts))
}

/**
 * A collection of arguments for invoking getSourceDynamodb.
 */
export interface GetSourceDynamodbOutputArgs {
    sourceId: pulumi.Input<string>;
}
