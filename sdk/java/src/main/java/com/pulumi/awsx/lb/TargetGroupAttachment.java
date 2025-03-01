// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.awsx.lb;

import com.pulumi.aws.lambda.Permission;
import com.pulumi.awsx.Utilities;
import com.pulumi.awsx.lb.TargetGroupAttachmentArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Attach an EC2 instance or Lambda to a Load Balancer. This will create required permissions if attaching to a Lambda Function.
 * 
 */
@ResourceType(type="awsx:lb:TargetGroupAttachment")
public class TargetGroupAttachment extends com.pulumi.resources.ComponentResource {
    /**
     * Auto-created Lambda permission, if targeting a Lambda function
     * 
     */
    @Export(name="lambdaPermission", type=Permission.class, parameters={})
    private Output</* @Nullable */ Permission> lambdaPermission;

    /**
     * @return Auto-created Lambda permission, if targeting a Lambda function
     * 
     */
    public Output<Optional<Permission>> lambdaPermission() {
        return Codegen.optional(this.lambdaPermission);
    }
    /**
     * Underlying Target Group Attachment resource
     * 
     */
    @Export(name="targetGroupAttachment", type=com.pulumi.aws.lb.TargetGroupAttachment.class, parameters={})
    private Output<com.pulumi.aws.lb.TargetGroupAttachment> targetGroupAttachment;

    /**
     * @return Underlying Target Group Attachment resource
     * 
     */
    public Output<com.pulumi.aws.lb.TargetGroupAttachment> targetGroupAttachment() {
        return this.targetGroupAttachment;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TargetGroupAttachment(String name) {
        this(name, TargetGroupAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TargetGroupAttachment(String name, @Nullable TargetGroupAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TargetGroupAttachment(String name, @Nullable TargetGroupAttachmentArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("awsx:lb:TargetGroupAttachment", name, args == null ? TargetGroupAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
