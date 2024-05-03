// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceStrava DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceStrava = airbyte.getSourceStrava({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceStrava(args: GetSourceStravaArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceStravaResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceStrava:getSourceStrava", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceStrava.
 */
export interface GetSourceStravaArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceStrava.
 */
export interface GetSourceStravaResult {
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
 * SourceStrava DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceStrava = airbyte.getSourceStrava({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceStravaOutput(args: GetSourceStravaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceStravaResult> {
    return pulumi.output(args).apply((a: any) => getSourceStrava(a, opts))
}

/**
 * A collection of arguments for invoking getSourceStrava.
 */
export interface GetSourceStravaOutputArgs {
    sourceId: pulumi.Input<string>;
}