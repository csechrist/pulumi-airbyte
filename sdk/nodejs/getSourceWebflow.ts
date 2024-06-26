// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceWebflow DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceWebflow = airbyte.getSourceWebflow({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceWebflow(args: GetSourceWebflowArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceWebflowResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceWebflow:getSourceWebflow", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceWebflow.
 */
export interface GetSourceWebflowArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceWebflow.
 */
export interface GetSourceWebflowResult {
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
 * SourceWebflow DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceWebflow = airbyte.getSourceWebflow({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceWebflowOutput(args: GetSourceWebflowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceWebflowResult> {
    return pulumi.output(args).apply((a: any) => getSourceWebflow(a, opts))
}

/**
 * A collection of arguments for invoking getSourceWebflow.
 */
export interface GetSourceWebflowOutputArgs {
    sourceId: pulumi.Input<string>;
}
