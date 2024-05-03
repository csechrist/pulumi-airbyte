// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceMarketo DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceMarketo = airbyte.getSourceMarketo({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceMarketo(args: GetSourceMarketoArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceMarketoResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceMarketo:getSourceMarketo", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceMarketo.
 */
export interface GetSourceMarketoArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceMarketo.
 */
export interface GetSourceMarketoResult {
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
 * SourceMarketo DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceMarketo = airbyte.getSourceMarketo({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceMarketoOutput(args: GetSourceMarketoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceMarketoResult> {
    return pulumi.output(args).apply((a: any) => getSourceMarketo(a, opts))
}

/**
 * A collection of arguments for invoking getSourceMarketo.
 */
export interface GetSourceMarketoOutputArgs {
    sourceId: pulumi.Input<string>;
}